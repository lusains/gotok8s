package adapter

import "testing"

func TestAdapter(t *testing.T) {
	audioPlayer := &AudioPlayer{}

	// 测试播放 mp3 文件
	audioPlayer.Play("mp3", "beyond the horizon.mp3")

	// 测试播放 mp4 文件
	audioPlayer.Play("mp4", "alone.mp4")

	// 测试播放 vlc 文件
	audioPlayer.Play("vlc", "far far away.vlc")

	// 测试播放不支持的格式
	audioPlayer.Play("avi", "mind me.avi")

	// 验证适配器的正确性
	mediaAdapter := NewMediaAdapter("mp4")
	if _, ok := mediaAdapter.advancedMusicPlayer.(*Mp4Player); !ok {
		t.Error("应该创建Mp4Player实例")
	}

	mediaAdapter = NewMediaAdapter("vlc")
	if _, ok := mediaAdapter.advancedMusicPlayer.(*VlcPlayer); !ok {
		t.Error("应该创建VlcPlayer实例")
	}
}
