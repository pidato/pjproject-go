package pjsua2

/*
#include <stdio.h>
const int PI_AUDIO_FRAME_MAX_PCM_BYTES = 640;
const int PI_AUDIO_FRAME_MAX_OPUS_BYTES = 1328;
typedef struct PiAudioFrameExt {
    void *encoder;
    uint64_t enqueuedAt;
    uint64_t processedCPU;
    uint16_t port;
    uint32_t cycle;
    uint64_t timestamp;
    int64_t frame_num;
    int64_t vad_cpu;
    int64_t opus_cpu;
    bool processed;
    bool dtx;
    int8_t vad;
    int16_t pcm_bytes;
    int16_t pcm_samples;
    int16_t opus_size;
    uint8_t pcm[PI_AUDIO_FRAME_MAX_PCM_BYTES];
    uint8_t opus[PI_AUDIO_FRAME_MAX_OPUS_BYTES];
} PiAudioFrameExt;
*/

type AudioFrame struct {
	Encoder      uintptr
	EnqueuedAt   uint64
	ProcessedCPU uint64
	Port         uint16
	Cycle        uint32
	Timestamp    uint64
	FrameNum     int64
	VadCPU       int64
	OpusCPU      int64
	Processed    bool
	DTX          bool
	Vad          int8
	PcmBytes     int16
	PcmSamples   int16
	OpusSize     int16
	Pcm          [640]byte
	Opus         [1328]byte
}
