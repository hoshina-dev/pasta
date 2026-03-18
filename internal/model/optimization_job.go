package model

type OptimizationJob struct {
	UUID                      string `json:"uuid"`
	SourceGLMURL              string `json:"source_glm_url"`
	DestGLMURL                string `json:"dest_glm_url"`
	WebhookURL                string `json:"webhook_url"`
	DracoCompressionLevel     int    `json:"draco_compression_level,omitempty"`
	DracoPositionQuantization int    `json:"draco_position_quantization,omitempty"`
	DracoTexcoordQuantization int    `json:"draco_texcoord_quantization,omitempty"`
	DracoNormalQuantization   int    `json:"draco_normal_quantization,omitempty"`
	DracoGenericQuantization  int    `json:"draco_generic_quantization,omitempty"`
}
