components {
  id: "deathfx"
  component: "/particles/death.particlefx"
  position {
    y: 40.0
    z: 0.1
  }
}
components {
  id: "script"
  component: "/hero/hero.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"hero\"\n"
  "mask: \"geometry\"\n"
  "mask: \"danger\"\n"
  "mask: \"pickup\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 30.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 74.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 1\n"
  "  }\n"
  "  data: 15.0\n"
  "  data: 30.0\n"
  "  data: 10.0\n"
  "  data: 15.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/def-runner/spineboy.spinescene\"\n"
  "default_animation: \"run\"\n"
  "skin: \"\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
  scale {
    x: 0.15
    y: 0.15
  }
}
