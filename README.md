# GoTracing Project

This is a ray tracing project which is based on golang.

The purpoes of the project is just for me to learn golang.

# What's the next

Now GoTracing can render some simple objects. But ray tracing is a complex system, I can't finish it in a short time. So I will make a Todo list.

1. ~~using configuration file to describe scene~~
2. ~~add local coordinate system for object~~
3. ~~rebuild the camra and view plane~~
4. enhance the shadow and reflection
5. multi kinds of lights
6. and so on

# Configuration

GoTracing can use configuration file to build scene. Here given a example file. Currently only sphere and rect can be described by configuration.

```yaml
---
main:
  width: 1280
  height: 720
  output: /home/example.jpg
  renderThreads: 4
tracer:
  kind: SimpleTracer
camra:
  position: 0,500,1000
  distance: 900
  u: 1,0,0
  v: 0,2,-1
  w: 0,1,2
  viewplane:
    width: 1280
    height: 720
    sample: 16
lights:
- name: light1
  kind: SimplePointLight
  args:
    position: 10000000,10000000,10000000
    color: 255,255,255,255
    ls: 1.0
objects:
- name: WhiteSphere
  kind: Sphere
  args:
    center: 0,120,0
    radius: 120
    localX: 1,0,0
    localY: 0,1,0
    localZ: 0,0,1
  material:
    kind: SpecularPhong
    args:
      ks: 0.3
      kd: 0.5
      exp: 3
      color: 255, 255, 255, 255
```


# Example

* some objects in one scene
![example](./img/example.jpg)

* object in different light color
![object in different light color](./img/differentLightColor.jpg)