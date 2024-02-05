components {
  id: "mainship2-health1"
  component: "/main/mainship2-health.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "mainship2-health"
  type: "sprite"
  data: "tile_set: \"/main/mainship2.atlas\"\n"
  "default_animation: \"effect_yellow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    x: 0.5
    y: 2.5
    z: 1.0
  }
}
