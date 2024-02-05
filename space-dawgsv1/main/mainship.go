components {
  id: "gravtrial.script"
  component: "/main/gravtrial.script.script"
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
  id: "mainship"
  type: "sprite"
  data: "tile_set: \"/main/spaceship.atlas\"\n"
  "default_animation: \"ship_L\"\n"
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
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.585185
    y: 0.711027
    z: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 5.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"mainship\"\n"
  "mask: \"hole\"\n"
  "mask: \"bullets2\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: -0.026035087\n"
  "      z: 0.0\n"
  "      w: 0.999661\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 15.0\n"
  "  data: 18.510332\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
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
  id: "factory"
  type: "factory"
  data: "prototype: \"/main/mainship-bullet.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
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
