package youtubeResponse

type YoutubeResponse struct {
	Attestation struct {
		PlayerAttestationRenderer struct {
			BotguardData struct {
				InterpreterSafeURL struct {
					PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				Program           string `json:"program"`
				ServerEnvironment int    `json:"serverEnvironment"`
			} `json:"botguardData"`
			Challenge string `json:"challenge"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	FrameworkUpdates struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Payload   struct {
					OfflineabilityEntity struct {
						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
						Key                     string `json:"key"`
					} `json:"offlineabilityEntity"`
				} `json:"payload"`
				Type string `json:"type"`
			} `json:"mutations"`
			Timestamp struct {
				Nanos   int    `json:"nanos"`
				Seconds string `json:"seconds"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			AvailableCountries []string `json:"availableCountries"`
			Category           string   `json:"category"`
			Description        struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			Embed struct {
				FlashSecureURL string `json:"flashSecureUrl"`
				FlashURL       string `json:"flashUrl"`
				Height         int    `json:"height"`
				IframeURL      string `json:"iframeUrl"`
				Width          int    `json:"width"`
			} `json:"embed"`
			ExternalChannelID string `json:"externalChannelId"`
			HasYpcMetadata    bool   `json:"hasYpcMetadata"`
			IsFamilySafe      bool   `json:"isFamilySafe"`
			IsUnlisted        bool   `json:"isUnlisted"`
			LengthSeconds     string `json:"lengthSeconds"`
			OwnerChannelName  string `json:"ownerChannelName"`
			OwnerProfileURL   string `json:"ownerProfileUrl"`
			PublishDate       string `json:"publishDate"`
			Thumbnail         struct {
				Thumbnails []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			Title struct {
				SimpleText string `json:"simpleText"`
			} `json:"title"`
			UploadDate string `json:"uploadDate"`
			ViewCount  string `json:"viewCount"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	PlayabilityStatus struct {
		ContextParams string `json:"contextParams"`
		Miniplayer    struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		PlayableInEmbed bool   `json:"playableInEmbed"`
		Status          string `json:"status"`
	} `json:"playabilityStatus"`
	PlaybackTracking struct {
		AtrURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"atrUrl"`
		PtrackingURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"ptrackingUrl"`
		QoeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"qoeUrl"`
		VideostatsDefaultFlushIntervalSeconds int `json:"videostatsDefaultFlushIntervalSeconds"`
		VideostatsDelayplayURL                struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsPlaybackURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsWatchtimeURL                  struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsWatchtimeUrl"`
	} `json:"playbackTracking"`
	PlayerConfig struct {
		AudioConfig struct {
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
		} `json:"audioConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		WebPlayerConfig struct {
			WebPlayerActionsPorting struct {
				AddToWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							APIURL   string `json:"apiUrl"`
							SendPost bool   `json:"sendPost"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						Actions []struct {
							Action       string `json:"action"`
							AddedVideoID string `json:"addedVideoId"`
						} `json:"actions"`
						PlaylistID string `json:"playlistId"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				GetSharePanelCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							APIURL   string `json:"apiUrl"`
							SendPost bool   `json:"sendPost"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WebPlayerShareEntityServiceEndpoint struct {
						SerializedShareEntity string `json:"serializedShareEntity"`
					} `json:"webPlayerShareEntityServiceEndpoint"`
				} `json:"getSharePanelCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							APIURL   string `json:"apiUrl"`
							SendPost bool   `json:"sendPost"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						Actions []struct {
							Action         string `json:"action"`
							RemovedVideoID string `json:"removedVideoId"`
						} `json:"actions"`
						PlaylistID string `json:"playlistId"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							APIURL   string `json:"apiUrl"`
							SendPost bool   `json:"sendPost"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					SubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							APIURL   string `json:"apiUrl"`
							SendPost bool   `json:"sendPost"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	ResponseContext struct {
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		ServiceTrackingParams []struct {
			Params []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
			Service string `json:"service"`
		} `json:"serviceTrackingParams"`
		VisitorData                     string `json:"visitorData"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec string `json:"spec"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	StreamingData struct {
		AdaptiveFormats []struct {
			ApproxDurationMs string `json:"approxDurationMs"`
			AverageBitrate   int    `json:"averageBitrate"`
			Bitrate          int    `json:"bitrate"`
			ContentLength    string `json:"contentLength"`
			Fps              int    `json:"fps,omitempty"`
			Height           int    `json:"height,omitempty"`
			IndexRange       struct {
				End   string `json:"end"`
				Start string `json:"start"`
			} `json:"indexRange"`
			InitRange struct {
				End   string `json:"end"`
				Start string `json:"start"`
			} `json:"initRange"`
			Itag           int    `json:"itag"`
			LastModified   string `json:"lastModified"`
			MimeType       string `json:"mimeType"`
			ProjectionType string `json:"projectionType"`
			Quality        string `json:"quality"`
			QualityLabel   string `json:"qualityLabel,omitempty"`
			URL            string `json:"url"`
			Width          int    `json:"width,omitempty"`
			ColorInfo      struct {
				MatrixCoefficients      string `json:"matrixCoefficients"`
				Primaries               string `json:"primaries"`
				TransferCharacteristics string `json:"transferCharacteristics"`
			} `json:"colorInfo,omitempty"`
			HighReplication bool    `json:"highReplication,omitempty"`
			AudioChannels   int     `json:"audioChannels,omitempty"`
			AudioQuality    string  `json:"audioQuality,omitempty"`
			AudioSampleRate string  `json:"audioSampleRate,omitempty"`
			LoudnessDb      float64 `json:"loudnessDb,omitempty"`
		} `json:"adaptiveFormats"`
		ExpiresInSeconds string `json:"expiresInSeconds"`
		Formats          []struct {
			ApproxDurationMs string `json:"approxDurationMs"`
			AudioChannels    int    `json:"audioChannels"`
			AudioQuality     string `json:"audioQuality"`
			AudioSampleRate  string `json:"audioSampleRate"`
			AverageBitrate   int    `json:"averageBitrate"`
			Bitrate          int    `json:"bitrate"`
			ContentLength    string `json:"contentLength"`
			Fps              int    `json:"fps"`
			Height           int    `json:"height"`
			Itag             int    `json:"itag"`
			LastModified     string `json:"lastModified"`
			MimeType         string `json:"mimeType"`
			ProjectionType   string `json:"projectionType"`
			Quality          string `json:"quality"`
			QualityLabel     string `json:"qualityLabel"`
			URL              string `json:"url"`
			Width            int    `json:"width"`
		} `json:"formats"`
	} `json:"streamingData"`
	TrackingParams string `json:"trackingParams"`
	VideoDetails   struct {
		AllowRatings      bool     `json:"allowRatings"`
		Author            string   `json:"author"`
		ChannelID         string   `json:"channelId"`
		IsCrawlable       bool     `json:"isCrawlable"`
		IsLiveContent     bool     `json:"isLiveContent"`
		IsOwnerViewing    bool     `json:"isOwnerViewing"`
		IsPrivate         bool     `json:"isPrivate"`
		IsUnpluggedCorpus bool     `json:"isUnpluggedCorpus"`
		Keywords          []string `json:"keywords"`
		LengthSeconds     string   `json:"lengthSeconds"`
		ShortDescription  string   `json:"shortDescription"`
		Thumbnail         struct {
			Thumbnails []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		Title     string `json:"title"`
		VideoID   string `json:"videoId"`
		ViewCount string `json:"viewCount"`
	} `json:"videoDetails"`
}
