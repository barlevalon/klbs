package tautulli

type GetActivityResponse struct {
	Response struct {
		Result  string `json:"result"`
		Message any    `json:"message"`
		Data    struct {
			StreamCount string `json:"stream_count"`
			Sessions    []struct {
				LastSeen                     any      `json:"last_seen"`
				AudioDecision                string   `json:"audio_decision"`
				VideoDecision                string   `json:"video_decision"`
				ProgressPercent              string   `json:"progress_percent"`
				QualityProfile               string   `json:"quality_profile"`
				SyncedVersionProfile         string   `json:"synced_version_profile"`
				OptimizedVersionProfile      string   `json:"optimized_version_profile"`
				User                         string   `json:"user"`
				StreamSubtitleDecision       string   `json:"stream_subtitle_decision"`
				SectionID                    string   `json:"section_id"`
				LibraryName                  string   `json:"library_name"`
				RatingKey                    string   `json:"rating_key"`
				ParentRatingKey              string   `json:"parent_rating_key"`
				GrandparentRatingKey         string   `json:"grandparent_rating_key"`
				Title                        string   `json:"title"`
				ParentTitle                  string   `json:"parent_title"`
				GrandparentTitle             string   `json:"grandparent_title"`
				OriginalTitle                string   `json:"original_title"`
				SortTitle                    string   `json:"sort_title"`
				EditionTitle                 string   `json:"edition_title"`
				MediaIndex                   string   `json:"media_index"`
				ParentMediaIndex             string   `json:"parent_media_index"`
				Studio                       string   `json:"studio"`
				ContentRating                string   `json:"content_rating"`
				Summary                      string   `json:"summary"`
				Tagline                      string   `json:"tagline"`
				Rating                       string   `json:"rating"`
				RatingImage                  string   `json:"rating_image"`
				AudienceRating               string   `json:"audience_rating"`
				AudienceRatingImage          string   `json:"audience_rating_image"`
				UserRating                   string   `json:"user_rating"`
				Duration                     string   `json:"duration"`
				Year                         string   `json:"year"`
				ParentYear                   string   `json:"parent_year"`
				GrandparentYear              string   `json:"grandparent_year"`
				Thumb                        string   `json:"thumb"`
				ParentThumb                  string   `json:"parent_thumb"`
				GrandparentThumb             string   `json:"grandparent_thumb"`
				Art                          string   `json:"art"`
				Banner                       string   `json:"banner"`
				OriginallyAvailableAt        string   `json:"originally_available_at"`
				AddedAt                      string   `json:"added_at"`
				UpdatedAt                    string   `json:"updated_at"`
				LastViewedAt                 string   `json:"last_viewed_at"`
				GUID                         string   `json:"guid"`
				ParentGUID                   string   `json:"parent_guid"`
				GrandparentGUID              string   `json:"grandparent_guid"`
				StreamSubtitleLanguageCode   string   `json:"stream_subtitle_language_code"`
				StreamSubtitleLanguage       string   `json:"stream_subtitle_language"`
				StreamSubtitleLocation       string   `json:"stream_subtitle_location"`
				StreamSubtitleFormat         string   `json:"stream_subtitle_format"`
				StreamSubtitleContainer      string   `json:"stream_subtitle_container"`
				StreamSubtitleCodec          string   `json:"stream_subtitle_codec"`
				StreamAudioDecision          string   `json:"stream_audio_decision"`
				StreamAudioLanguageCode      string   `json:"stream_audio_language_code"`
				StreamAudioLanguage          string   `json:"stream_audio_language"`
				SubtitleLanguageCode         string   `json:"subtitle_language_code"`
				FullTitle                    string   `json:"full_title"`
				StreamAudioSampleRate        string   `json:"stream_audio_sample_rate"`
				StreamAudioCodec             string   `json:"stream_audio_codec"`
				ID                           string   `json:"id"`
				Container                    string   `json:"container"`
				Bitrate                      string   `json:"bitrate"`
				Height                       string   `json:"height"`
				Width                        string   `json:"width"`
				AspectRatio                  string   `json:"aspect_ratio"`
				VideoCodec                   string   `json:"video_codec"`
				VideoResolution              string   `json:"video_resolution"`
				VideoFullResolution          string   `json:"video_full_resolution"`
				VideoFramerate               string   `json:"video_framerate"`
				VideoProfile                 string   `json:"video_profile"`
				AudioCodec                   string   `json:"audio_codec"`
				AudioChannels                string   `json:"audio_channels"`
				AudioChannelLayout           string   `json:"audio_channel_layout"`
				AudioProfile                 string   `json:"audio_profile"`
				StreamAudioChannelLayout     string   `json:"stream_audio_channel_layout"`
				ChannelCallSign              string   `json:"channel_call_sign"`
				ChannelIdentifier            string   `json:"channel_identifier"`
				ChannelThumb                 string   `json:"channel_thumb"`
				File                         string   `json:"file"`
				FileSize                     string   `json:"file_size"`
				StreamAudioChannels          string   `json:"stream_audio_channels"`
				StreamAudioBitrateMode       string   `json:"stream_audio_bitrate_mode"`
				Type                         string   `json:"type"`
				VideoCodecLevel              string   `json:"video_codec_level"`
				VideoBitrate                 string   `json:"video_bitrate"`
				VideoBitDepth                string   `json:"video_bit_depth"`
				VideoChromaSubsampling       string   `json:"video_chroma_subsampling"`
				VideoColorPrimaries          string   `json:"video_color_primaries"`
				VideoColorRange              string   `json:"video_color_range"`
				VideoColorSpace              string   `json:"video_color_space"`
				VideoColorTrc                string   `json:"video_color_trc"`
				VideoDynamicRange            string   `json:"video_dynamic_range"`
				VideoFrameRate               string   `json:"video_frame_rate"`
				VideoRefFrames               string   `json:"video_ref_frames"`
				VideoHeight                  string   `json:"video_height"`
				VideoWidth                   string   `json:"video_width"`
				VideoLanguage                string   `json:"video_language"`
				VideoLanguageCode            string   `json:"video_language_code"`
				VideoScanType                string   `json:"video_scan_type"`
				AudioBitrate                 string   `json:"audio_bitrate"`
				AudioBitrateMode             string   `json:"audio_bitrate_mode"`
				AudioSampleRate              string   `json:"audio_sample_rate"`
				AudioLanguage                string   `json:"audio_language"`
				AudioLanguageCode            string   `json:"audio_language_code"`
				SubtitleCodec                string   `json:"subtitle_codec"`
				SubtitleContainer            string   `json:"subtitle_container"`
				SubtitleFormat               string   `json:"subtitle_format"`
				StreamAudioBitrate           string   `json:"stream_audio_bitrate"`
				SubtitleLocation             string   `json:"subtitle_location"`
				SubtitleLanguage             string   `json:"subtitle_language"`
				StreamAudioChannelLayout0    string   `json:"stream_audio_channel_layout_"`
				State                        string   `json:"state"`
				StreamVideoLanguageCode      string   `json:"stream_video_language_code"`
				Username                     string   `json:"username"`
				FriendlyName                 string   `json:"friendly_name"`
				UserThumb                    string   `json:"user_thumb"`
				Email                        string   `json:"email"`
				StreamVideoDecision          string   `json:"stream_video_decision"`
				StreamVideoScanType          string   `json:"stream_video_scan_type"`
				StreamVideoLanguage          string   `json:"stream_video_language"`
				StreamVideoRefFrames         string   `json:"stream_video_ref_frames"`
				StreamVideoWidth             string   `json:"stream_video_width"`
				StreamVideoHeight            string   `json:"stream_video_height"`
				StreamVideoDynamicRange      string   `json:"stream_video_dynamic_range"`
				StreamVideoColorTrc          string   `json:"stream_video_color_trc"`
				StreamVideoColorSpace        string   `json:"stream_video_color_space"`
				StreamVideoColorRange        string   `json:"stream_video_color_range"`
				MediaType                    string   `json:"media_type"`
				IPAddress                    string   `json:"ip_address"`
				IPAddressPublic              string   `json:"ip_address_public"`
				Device                       string   `json:"device"`
				Platform                     string   `json:"platform"`
				PlatformName                 string   `json:"platform_name"`
				PlatformVersion              string   `json:"platform_version"`
				Product                      string   `json:"product"`
				ProductVersion               string   `json:"product_version"`
				Profile                      string   `json:"profile"`
				Player                       string   `json:"player"`
				MachineID                    string   `json:"machine_id"`
				StreamVideoColorPrimaries    string   `json:"stream_video_color_primaries"`
				StreamVideoCodecLevel        string   `json:"stream_video_codec_level"`
				StreamVideoCodec             string   `json:"stream_video_codec"`
				StreamVideoChromaSubsampling string   `json:"stream_video_chroma_subsampling"`
				SessionID                    string   `json:"session_id"`
				Bandwidth                    string   `json:"bandwidth"`
				Location                     string   `json:"location"`
				TranscodeKey                 string   `json:"transcode_key"`
				StreamVideoBitDepth          string   `json:"stream_video_bit_depth"`
				StreamVideoBitrate           string   `json:"stream_video_bitrate"`
				TranscodeSpeed               string   `json:"transcode_speed"`
				TranscodeAudioChannels       string   `json:"transcode_audio_channels"`
				TranscodeAudioCodec          string   `json:"transcode_audio_codec"`
				TranscodeVideoCodec          string   `json:"transcode_video_codec"`
				TranscodeWidth               string   `json:"transcode_width"`
				TranscodeHeight              string   `json:"transcode_height"`
				TranscodeContainer           string   `json:"transcode_container"`
				TranscodeProtocol            string   `json:"transcode_protocol"`
				StreamVideoFullResolution    string   `json:"stream_video_full_resolution"`
				ContainerDecision            string   `json:"container_decision"`
				TranscodeDecision            string   `json:"transcode_decision"`
				TranscodeHwDecode            string   `json:"transcode_hw_decode"`
				TranscodeHwDecodeTitle       string   `json:"transcode_hw_decode_title"`
				TranscodeHwEncode            string   `json:"transcode_hw_encode"`
				TranscodeHwEncodeTitle       string   `json:"transcode_hw_encode_title"`
				BifThumb                     string   `json:"bif_thumb"`
				SessionKey                   string   `json:"session_key"`
				ViewOffset                   string   `json:"view_offset"`
				SubtitleDecision             string   `json:"subtitle_decision"`
				Throttled                    string   `json:"throttled"`
				LiveUUID                     string   `json:"live_uuid"`
				OptimizedVersionTitle        string   `json:"optimized_version_title"`
				StreamContainer              string   `json:"stream_container"`
				StreamBitrate                string   `json:"stream_bitrate"`
				StreamAspectRatio            string   `json:"stream_aspect_ratio"`
				StreamVideoFramerate         string   `json:"stream_video_framerate"`
				StreamVideoResolution        string   `json:"stream_video_resolution"`
				StreamDuration               string   `json:"stream_duration"`
				StreamContainerDecision      string   `json:"stream_container_decision"`
				SharedLibraries              []string `json:"shared_libraries"`
				Directors                    []string `json:"directors"`
				Writers                      []string `json:"writers"`
				Actors                       []string `json:"actors"`
				Genres                       []string `json:"genres"`
				Labels                       []any    `json:"labels"`
				Collections                  []any    `json:"collections"`
				Guids                        []string `json:"guids"`
				Markers                      []struct {
					Type            string `json:"type"`
					ID              int    `json:"id"`
					StartTimeOffset int    `json:"start_time_offset"`
					EndTimeOffset   int    `json:"end_time_offset"`
					First           bool   `json:"first"`
					Final           bool   `json:"final"`
				} `json:"markers"`
				ParentGuids                 []string `json:"parent_guids"`
				GrandparentGuids            []string `json:"grandparent_guids"`
				Relayed                     int      `json:"relayed"`
				OptimizedVersion            int      `json:"optimized_version"`
				RowID                       int      `json:"row_id"`
				TranscodeHwEncoding         int      `json:"transcode_hw_encoding"`
				AllowGuest                  int      `json:"allow_guest"`
				IsAdmin                     int      `json:"is_admin"`
				KeepHistory                 int      `json:"keep_history"`
				DoNotify                    int      `json:"do_notify"`
				IsRestricted                int      `json:"is_restricted"`
				IsAllowSync                 int      `json:"is_allow_sync"`
				TranscodeHwDecoding         int      `json:"transcode_hw_decoding"`
				ChannelStream               int      `json:"channel_stream"`
				DeletedUser                 int      `json:"deleted_user"`
				IsActive                    int      `json:"is_active"`
				SubtitleForced              int      `json:"subtitle_forced"`
				Selected                    int      `json:"selected"`
				Indexes                     int      `json:"indexes"`
				Local                       int      `json:"local"`
				Live                        int      `json:"live"`
				ChildrenCount               int      `json:"children_count"`
				Secure                      int      `json:"secure"`
				TranscodeThrottled          int      `json:"transcode_throttled"`
				TranscodeProgress           int      `json:"transcode_progress"`
				TranscodeMinOffsetAvailable int      `json:"transcode_min_offset_available"`
				TranscodeMaxOffsetAvailable int      `json:"transcode_max_offset_available"`
				TranscodeHwRequested        int      `json:"transcode_hw_requested"`
				Subtitles                   int      `json:"subtitles"`
				StreamSubtitleForced        int      `json:"stream_subtitle_forced"`
				TranscodeHwFullPipeline     int      `json:"transcode_hw_full_pipeline"`
				IsHomeUser                  int      `json:"is_home_user"`
				SyncedVersion               int      `json:"synced_version"`
				UserID                      int      `json:"user_id"`
				StreamSubtitleTransient     int      `json:"stream_subtitle_transient"`
			} `json:"sessions"`
			StreamCountDirectPlay   int `json:"stream_count_direct_play"`
			StreamCountDirectStream int `json:"stream_count_direct_stream"`
			StreamCountTranscode    int `json:"stream_count_transcode"`
			TotalBandwidth          int `json:"total_bandwidth"`
			LanBandwidth            int `json:"lan_bandwidth"`
			WanBandwidth            int `json:"wan_bandwidth"`
		} `json:"data"`
	} `json:"response"`
}
