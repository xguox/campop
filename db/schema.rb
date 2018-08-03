# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20180727035729) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "cameras", force: :cascade do |t|
    t.string "product_type"
    t.string "sensor_type"
    t.string "release_date"
    t.string "discontinued"
    t.string "image_sensor_size"
    t.string "total_pixels"
    t.string "effective_pixels"
    t.string "image_processor"
    t.string "static_image_format"
    t.string "static_image_size"
    t.string "frame_ratio"
    t.string "movie_format"
    t.string "movie_size"
    t.string "audio_format"
    t.string "dust_removal_system"
    t.string "lens_structure"
    t.string "aspherical_lenses"
    t.string "focal_length"
    t.string "eq_35mm_frame_focal_length"
    t.string "maximum_aperture"
    t.string "optical_zoom"
    t.string "recent_focusing_distance"
    t.string "focus_mode"
    t.string "lens_mount"
    t.string "compatible_lens"
    t.string "focus_system_type"
    t.string "focus"
    t.string "metering_system"
    t.string "metering_mode"
    t.string "af_assist_light"
    t.string "exposure_mode"
    t.string "shooting_mode"
    t.string "film_simulation_mode"
    t.string "advanced_filter"
    t.string "shutter_type"
    t.string "shutter_speed"
    t.string "flash_sync_speed"
    t.string "exposure_compensation"
    t.string "iso_sensitivity"
    t.string "white_balance"
    t.string "timer"
    t.string "continuous_shot"
    t.string "viewfinder"
    t.string "viewing_range"
    t.string "viewing_magnification"
    t.string "diopter_adjustment_range"
    t.string "focus_screen"
    t.string "reflector"
    t.string "deep_depth_preview"
    t.string "lcd_type"
    t.string "lcd_size"
    t.string "lcd_pixels"
    t.string "field_of_view"
    t.string "lcd_characteristics"
    t.string "storage_card"
    t.string "data_interface"
    t.string "video_audio_interface"
    t.string "wireless_function"
    t.string "battery"
    t.string "battery_life"
    t.string "colour"
    t.string "size"
    t.string "weight"
    t.string "operating_temperature"
    t.string "working_humidity"
    t.string "annex"
  end

  create_table "lenses", force: :cascade do |t|
    t.string "lens_type"
    t.string "release_date"
    t.string "focal_length"
    t.string "eq_35mm_frame_focal_length"
    t.string "maximum_aperture"
    t.string "minimum_aperture"
    t.string "viewing_angle"
    t.string "lens_structure"
    t.string "aspherical_lens"
    t.string "low_dispersion_lens"
    t.string "aperture_blade"
    t.string "recent_focus_distance"
    t.string "maximum_magnification"
    t.string "filter_aperture"
    t.string "lens_mount"
    t.string "image_stabilizer"
    t.string "colour"
    t.string "size"
    t.string "weight"
    t.string "annex"
  end

end
