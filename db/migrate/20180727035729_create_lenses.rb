class CreateLenses < ActiveRecord::Migration[5.1]
  def change
    create_table :lenses do |t|
      t.string :lens_type
      t.string :release_date
      t.string :focal_length
      t.string :eq_35mm_frame_focal_length
      t.string :maximum_aperture
      t.string :minimum_aperture
      t.string :viewing_angle
      t.string :lens_structure
      t.string :aspherical_lens
      t.string :low_dispersion_lens
      t.string :aperture_blade
      t.string :recent_focus_distance
      t.string :maximum_magnification
      t.string :filter_aperture
      t.string :lens_mount
      t.string :image_stabilizer
      t.string :colour
      t.string :size
      t.string :weight
      t.string :annex
    end
  end
end
