//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package tool

import (
	"encoding/json"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Media generation tools. The server runs these itself, calling the AI
// provider named in the config, so no worker is involved. The model decides
// when to call one and supplies the dynamic parameters; the provider and
// model name are fixed here.
//
// The default input schemas are the Python SDK's, field for field, so an
// agent declared in either SDK compiles to the same tool. Override one with
// WithInputSchema; add static generation parameters with WithConfig, as
// Python's **defaults do.

// Image generates an image with the given provider and model, for example
// "openai" and "dall-e-3".
func Image(name, description, llmProvider, model string, opts ...Option) ai.ToolDef {
	return mediaTool(name, description, ai.ToolTypeGenerateImage, "GENERATE_IMAGE",
		llmProvider, model, imageSchema(), opts)
}

// Audio generates speech with the given provider and model, for example
// "openai" and "tts-1".
func Audio(name, description, llmProvider, model string, opts ...Option) ai.ToolDef {
	return mediaTool(name, description, ai.ToolTypeGenerateAudio, "GENERATE_AUDIO",
		llmProvider, model, audioSchema(), opts)
}

// Video generates a video with the given provider and model, for example
// "openai" and "sora".
func Video(name, description, llmProvider, model string, opts ...Option) ai.ToolDef {
	return mediaTool(name, description, ai.ToolTypeGenerateVideo, "GENERATE_VIDEO",
		llmProvider, model, videoSchema(), opts)
}

// PDF renders Markdown to a PDF. It needs no provider.
func PDF(name, description string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: pdfSchema(),
		ToolType:    ai.ToolTypeGeneratePDF,
		Config:      map[string]any{"taskType": "GENERATE_PDF"},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

func mediaTool(name, description string, toolType ai.ToolType, taskType, llmProvider, model string,
	schema map[string]any, opts []Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: schema,
		ToolType:    toolType,
		Config: map[string]any{
			"taskType":    taskType,
			"llmProvider": llmProvider,
			"model":       model,
		},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// prop builds one JSON Schema property. A nil def means no default.
func prop(typ, description string, def any) map[string]any {
	p := map[string]any{"type": typ, "description": description}
	if def != nil {
		p["default"] = def
	}
	return p
}

func objectSchema(required []string, properties map[string]any) map[string]any {
	return map[string]any{"type": "object", "properties": properties, "required": required}
}

func imageSchema() map[string]any {
	return objectSchema([]string{"prompt"}, map[string]any{
		"prompt":       prop("string", "Text description of the image to generate.", nil),
		"style":        prop("string", "Image style: 'vivid' or 'natural'.", nil),
		"size":         prop("string", "Image size (e.g. '1024x1024'). Alternative to width/height.", nil),
		"width":        prop("integer", "Image width in pixels.", 1024),
		"height":       prop("integer", "Image height in pixels.", 1024),
		"n":            prop("integer", "Number of images to generate.", 1),
		"outputFormat": prop("string", "Output format: 'png', 'jpg', or 'webp'.", "png"),
		"weight":       prop("number", "Image weight parameter.", nil),
	})
}

func audioSchema() map[string]any {
	voice := prop("string", "Voice to use.", "alloy")
	voice["enum"] = []string{"alloy", "echo", "fable", "onyx", "nova", "shimmer"}
	return objectSchema([]string{"text"}, map[string]any{
		"text":  prop("string", "Text to convert to speech.", nil),
		"voice": voice,
		// json.Number keeps the literal "1.0": Python writes the default as
		// 1.0, and a float64 would encode as 1, which the server stores as a
		// different JSON number type and the recorder treats as a different
		// tool schema.
		"speed":          prop("number", "Speech speed multiplier (0.25 to 4.0).", json.Number("1.0")),
		"responseFormat": prop("string", "Audio format: 'mp3', 'wav', 'opus', 'aac', or 'flac'.", "mp3"),
		"n":              prop("integer", "Number of audio outputs to generate.", 1),
	})
}

func videoSchema() map[string]any {
	return objectSchema([]string{"prompt"}, map[string]any{
		"prompt":             prop("string", "Text description of the video scene.", nil),
		"negativePrompt":     prop("string", "Description of what to exclude from the video.", nil),
		"duration":           prop("integer", "Video duration in seconds.", 5),
		"maxDurationSeconds": prop("integer", "Maximum duration ceiling in seconds.", nil),
		"width":              prop("integer", "Video width in pixels.", 1280),
		"height":             prop("integer", "Video height in pixels.", 720),
		"size":               prop("string", "Video size specification (e.g. '1280x720').", nil),
		"resolution":         prop("string", "Quality level (e.g. '720p', '1080p').", nil),
		"aspectRatio":        prop("string", "Aspect ratio (e.g. '16:9', '1:1').", nil),
		"fps":                prop("integer", "Frames per second.", 24),
		"style":              prop("string", "Video style (e.g. 'cinematic', 'natural').", nil),
		"motion":             prop("string", "Movement intensity (e.g. 'slow', 'normal', 'extreme').", nil),
		"guidanceScale":      prop("number", "Prompt adherence strength (1.0 to 20.0).", nil),
		"seed":               prop("integer", "Seed for reproducibility.", nil),
		"inputImage":         prop("string", "Base64-encoded or URL image for image-to-video generation.", nil),
		"generateAudio":      prop("boolean", "Whether to generate audio with the video.", nil),
		"personGeneration":   prop("string", "Controls for human figure generation.", nil),
		"maxCostDollars":     prop("number", "Maximum cost limit in dollars.", nil),
		"outputFormat":       prop("string", "Video format (e.g. 'mp4').", "mp4"),
		"n":                  prop("integer", "Number of videos to generate.", 1),
	})
}

func pdfSchema() map[string]any {
	return objectSchema([]string{"markdown"}, map[string]any{
		"markdown":     prop("string", "Markdown text to convert to PDF.", nil),
		"pageSize":     prop("string", "Page size: A4, LETTER, LEGAL, A3, or A5.", "A4"),
		"theme":        prop("string", "Style preset: 'default' or 'compact'.", "default"),
		"baseFontSize": prop("number", "Base font size in points.", 11),
	})
}
