#version 460 core

#include<flutter/runtime_effect.glsl>

out vec4 fragColor;

uniform float time;
uniform vec2 resolution;

const float PI = 3.14159265359;

// 六边形网格函数
vec4 hexCoord(vec2 uv) {
    const vec2 r = vec2(1.0, sqrt(3.0));
    const vec2 h = r * 0.5;
    
    vec2 a = mod(uv, r) - h;
    vec2 b = mod(uv - h, r) - h;
    
    vec2 gv = length(a) < length(b) ? a : b;
    vec2 id = uv - gv;
    
    // 计算六边形的边界，使用线性插值确保边界共享
    float x = abs(gv.x);
    float y = abs(gv.y);
    float d = max(x * 0.866025 + y * 0.5, y) * 0.95; // 调整边界计算，使边缘更紧密
    
    return vec4(gv, id);
}

// 随机函数
float random(vec2 st) {
    return fract(sin(dot(st.xy, vec2(12.9898,78.233))) * 43758.5453123);
}

void main() {
    vec2 uv = (gl_FragCoord.xy * 2.0 - resolution.xy) / min(resolution.x, resolution.y);
    uv *= 4.0; // 增大蜂巢大小
    
    vec4 hc = hexCoord(uv);
    vec2 id = hc.zw;
    
    // 优化边框效果，使边缘共享
    float x = abs(hc.x);
    float y = abs(hc.y);
    float d = max(x * 0.866025 + y * 0.5, y);
    float cell = 1.0 - smoothstep(0.42, 0.45, d); // 调整边框宽度和过渡
    
    // 改进随机激活效果
    float t = time * 0.3; // 减慢动画速度
    float isActive = step(0.8, random(id + floor(t))); // 提高激活阈值
    
    // 增强视觉效果
    vec3 baseColor = vec3(0.1, 0.12, 0.15); // 深色背景
    vec3 borderColor = vec3(0.3, 0.35, 0.4); // 边框颜色
    vec3 activeColor = vec3(0.4, 0.45, 0.5); // 激活颜色
    
    vec3 col = mix(baseColor, borderColor, cell);
    col = mix(col, activeColor, isActive * cell);
    
    fragColor = vec4(col, 1.0);
}