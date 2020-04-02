# GoTracing Project

This is a ray tracing project which is based on golang.

The purpoes of the project is just for me to learn golang.

# What's the next

Now GoTracing can render some simple objects. But ray tracing is a complex system, I can't finish it in a short time. So I will make a Todo list.

1. ~~using configuration file to describe scene~~
2. add local coordinate system for object
3. rebuild the camra and view plane
4. enhance the shadow and reflection
5. multi kinds of color
6. and so on

# Configuration

GoTracing can use configuration file to build scene. Here given a example file. Currently only sphere and rect can be described by configuration.

```yaml
---
main:
  width: 800
  height: 800
  output: /home/example.jpg
  renderThreads: 4
camra:
  position: 0,0,0
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