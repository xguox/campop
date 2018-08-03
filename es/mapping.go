package es

const CameraIndexName = "cameras-es"
const CameraTypeName = "camera"
const CameraMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"camera":{
			"properties":{
				"id": {
					"type": "integer"
				},
				"name": {
					"type": "text"
				},
				"product_type": {
					"type": "text"
				},
				"sensor_type": {
					"type": "text"
				},
				"release_date": {
					"type": "text"
				},
				"discontinued": {
					"type": "text"
				},
				"image_sensor_size": {
					"type": "text"
				},
				"total_pixels": {
					"type": "text"
				},
				"effective_pixels": {
					"type": "text"
				},
				"image_processor": {
					"type": "text"
				},
				"static_image_format": {
					"type": "text"
				},
				"static_image_size": {
					"type": "text"
				},
				"frame_ratio": {
					"type": "text"
				},
				"movie_format": {
					"type": "text"
				},
				"movie_size": {
					"type": "text"
				},
				"audio_format": {
					"type": "text"
				},
				"dust_removal_system": {
					"type": "text"
				},
				"lens_structure": {
					"type": "text"
				},
				"aspherical_lenses": {
					"type": "text"
				},
				"focal_length": {
					"type": "text"
				},
				"eq_full_frame_focal_length": {
					"type": "text"
				},
				"maximum_aperture": {
					"type": "text"
				},
				"optical_zoom": {
					"type": "text"
				},
				"recent_focusing_distance": {
					"type": "text"
				},
				"focus_mode": {
					"type": "text"
				},
				"lens_mount": {
					"type": "text"
				},
				"compatible_lens": {
					"type": "text"
				},
				"focus_system_type": {
					"type": "text"
				},
				"focus": {
					"type": "text"
				},
				"metering_system": {
					"type": "text"
				},
				"metering_mode": {
					"type": "text"
				},
				"af_assist_light": {
					"type": "text"
				},
				"exposure_mode": {
					"type": "text"
				},
				"shooting_mode": {
					"type": "text"
				},
				"film_simulation_mode": {
					"type": "text"
				},
				"advanced_filter": {
					"type": "text"
				},
				"shutter_type": {
					"type": "text"
				},
				"shutter_speed": {
					"type": "text"
				},
				"flash_sync_speed": {
					"type": "text"
				},
				"exposure_compensation": {
					"type": "text"
				},
				"iso_sensitivity": {
					"type": "text"
				},
				"white_balance": {
					"type": "text"
				},
				"timer": {
					"type": "text"
				},
				"continuous_shot": {
					"type": "text"
				},
				"viewfinder": {
					"type": "text"
				},
				"viewing_range": {
					"type": "text"
				},
				"viewing_magnification": {
					"type": "text"
				},
				"diopter_adjustment_range": {
					"type": "text"
				},
				"focus_screen": {
					"type": "text"
				},
				"reflector": {
					"type": "text"
				},
				"deep_depth_preview": {
					"type": "text"
				},
				"lcd_type": {
					"type": "text"
				},
				"lcd_size": {
					"type": "text"
				},
				"lcd_pixels": {
					"type": "text"
				},
				"field_of_view": {
					"type": "text"
				},
				"lcd_characteristics": {
					"type": "text"
				},
				"storage_card": {
					"type": "text"
				},
				"data_interface": {
					"type": "text"
				},
				"video_audio_interface": {
					"type": "text"
				},
				"wireless_function": {
					"type": "text"
				},
				"battery": {
					"type": "text"
				},
				"battery_life": {
					"type": "text"
				},
				"colour": {
					"type": "text"
				},
				"size": {
					"type": "text"
				},
				"weight": {
					"type": "text"
				},
				"operating_temperature": {
					"type": "text"
				},
				"working_humidity": {
					"type": "text"
				},
				"annex": {
					"type": "text"
				}
			}
		}
	}
}`

const LensIndexName = "lenses-es"
const LensTypeName = "lens"
const LensMapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"lens":{
			"properties":{
				"id": {
					"type": "integer"
				},
				"name": {
					"type": "text"
				},
				"lens_type": {
					"type": "text"
				},
				"release_date": {
					"type": "text"
				},
				"focal_length": {
					"type": "text"
				},
				"eq_full_frame_focal_length": {
					"type": "text"
				},
				"maximum_aperture": {
					"type": "text"
				},
				"minimum_aperture": {
					"type": "text"
				},
				"viewing_angle": {
					"type": "text"
				},
				"lens_structure": {
					"type": "text"
				},
				"aspherical_lens": {
					"type": "text"
				},
				"low_dispersion_lens": {
					"type": "text"
				},
				"aperture_blade": {
					"type": "text"
				},
				"recent_focus_distance": {
					"type": "text"
				},
				"maximum_magnification": {
					"type": "text"
				},
				"filter_aperture": {
					"type": "text"
				},
				"lens_mount": {
					"type": "text"
				},
				"image_stabilizer": {
					"type": "text"
				},
				"colour": {
					"type": "text"
				},
				"size": {
					"type": "text"
				},
				"weight": {
					"type": "text"
				},
				"annex": {
					"type": "text"
				},
				"remarks": {
					"type": "text"
				}
			}
		}
	}
}`
