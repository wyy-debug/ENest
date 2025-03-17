cmake_minimum_required(VERSION 3.14...3.25)
set(PROJECT_NAME "win_zoom")
project(${PROJECT_NAME} LANGUAGES CXX)


set(PLUGIN_NAME "win_zoom_plugin")

link_directories("${PROJECT_SOURCE_DIR}/lib/") 

list(APPEND PLUGIN_SOURCES
  "win_zoom_plugin.cpp"
  "win_zoom_plugin.h"
)

add_library(${PLUGIN_NAME} SHARED
  "include/win_zoom/win_zoom_plugin_c_api.h"
  "win_zoom_plugin_c_api.cpp"
  ${PLUGIN_SOURCES}
)

apply_standard_settings(${PLUGIN_NAME})

set_target_properties(${PLUGIN_NAME} PROPERTIES
  CXX_VISIBILITY_PRESET hidden)
target_compile_definitions(${PLUGIN_NAME} PRIVATE FLUTTER_PLUGIN_IMPL)
target_include_directories(${PLUGIN_NAME} INTERFACE
  "${CMAKE_CURRENT_SOURCE_DIR}/include")

target_link_libraries(${PLUGIN_NAME} PRIVATE flutter flutter_wrapper_plugin videosdk)

set(win_zoom_bundled_libraries
  "${PROJECT_SOURCE_DIR}/bin/"
  PARENT_SCOPE
)
