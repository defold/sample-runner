components {
  id: "script"
  component: "/level/objects/platform.script"
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/level/objects/coin.go\"\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1.0539199\n"
  "      y: 21.369343\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 360.0\n"
  "  data: 40.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "danger_edges"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.7592307\n"
  "      y: -37.69981\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 380.0\n"
  "  data: 40.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -189.78683
    y: -0.6551108
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 190.17743
    y: -0.6551108
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -385.325
    y: -6.4391923
    z: -0.1
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -271.49286
    y: -76.12615
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -112.22625
    y: -70.68114
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite5"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 79.710434
    y: -72.04239
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite6"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 243.0608
    y: -67.66497
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite7"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/def-runner/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 376.21857
    y: -0.31878883
    z: -0.1
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
