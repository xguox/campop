class CreateCameras < ActiveRecord::Migration[5.1]
  def change
    create_table :cameras do |t|
      t.string :product_type
      t.string :sensor_type
      t.string :release_date
      t.string :discontinued
      t.string :image_sensor_size
      t.string :total_pixels
      t.string :effective_pixels
      t.string :image_processor
      t.string :static_image_format
      t.string :static_image_size
      t.string :frame_ratio
      t.string :movie_format
      t.string :movie_size
      t.string :audio_format
      t.string :dust_removal_system
      t.string :lens_structure
      t.string :aspherical_lenses
      t.string :focal_length
      t.string :eq_35mm_frame_focal_length
      t.string :maximum_aperture
      t.string :optical_zoom
      t.string :recent_focusing_distance
      t.string :focus_mode
      t.string :lens_mount
      t.string :compatible_lens
      t.string :focus_system_type
      t.string :focus
      t.string :metering_system
      t.string :metering_mode
      t.string :af_assist_light
      t.string :exposure_mode
      t.string :shooting_mode
      t.string :film_simulation_mode
      t.string :advanced_filter
      t.string :shutter_type
      t.string :shutter_speed
      t.string :flash_sync_speed
      t.string :exposure_compensation
      t.string :iso_sensitivity
      t.string :white_balance
      t.string :timer
      t.string :continuous_shot
      t.string :viewfinder
      t.string :viewing_range
      t.string :viewing_magnification
      t.string :diopter_adjustment_range
      t.string :focus_screen
      t.string :reflector
      t.string :deep_depth_preview
      t.string :lcd_type
      t.string :lcd_size
      t.string :lcd_pixels
      t.string :field_of_view
      t.string :lcd_characteristics
      t.string :storage_card
      t.string :data_interface
      t.string :video_audio_interface
      t.string :wireless_function
      t.string :battery
      t.string :battery_life
      t.string :colour
      t.string :size
      t.string :weight
      t.string :operating_temperature
      t.string :working_humidity
      t.string :annex
    end
  end
end
