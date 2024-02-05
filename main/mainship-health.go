components {
  id: "mainship-health1"
  component: "/main/mainship-health.script"
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
  id: "mainship-health"
  type: "sprite"
  data: "tile_set: \"/main/spaceship.atlas\"\n"
  "default_animation: \"effect_pink\"\n"
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
    z: -0.70954394
    w: 0.7046612
  }
  scale {
    x: 0.5
    y: 2.5
    z: 1.0
  }
}
