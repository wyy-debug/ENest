
<script setup>
import { onMounted, onUnmounted } from 'vue'
import uitoolkit from "@zoom/videosdk-ui-toolkit";
import '@zoom/videosdk-ui-toolkit/dist/videosdk-ui-toolkit.css'

var sessionContainer
var authEndpoint = '/zoom-auth'
var config = {
    videoSDKJWT: '',
    sessionName: 'test',
    userName: 'Vue.js',
    sessionPasscode: '123',
    deviceOptions: {
      video: true,
      audio: true
    },
    preloadVideo: false,
    rootElement: '#sessionContainer',
    featuresOptions: {
      video: true,
      audio: true,
      share: true,
      chat: true,
      virtualBackground: {
          enable: true,
          virtualBackgrounds: [
              {
                  url: 'https://images.unsplash.com/photo-1715490187538-30a365fa05bd?q=80&w=1945&auto=format&fit=crop'
              },
          ],
      },
  },
};
var role = 1

onMounted(() => {
  sessionContainer = document.getElementById('sessionContainer')
  if (!sessionContainer) {
    console.error('Session container not found')
    return
  }
})

onUnmounted(() => {
  if (sessionContainer) {
    uitoolkit.closeSession(sessionContainer)
  }
})

function getVideoSDKJWT() {
  if (!sessionContainer) {
    console.error('Session container not initialized')
    return
  }

  document.getElementById('join-flow').style.display = 'none'

  fetch(authEndpoint, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
        sessionName:  config.sessionName,
        role: role,
    })
  }).then((response) => {
      return response.json()
  }).then((data) => {
    if(data.signature) {
      console.log(data.signature)
      config.videoSDKJWT = data.signature
      joinSession()
    } else {
      console.log(data)
      document.getElementById('join-flow').style.display = 'block'
    }
  }).catch((error) => {
      console.error('Failed to get JWT:', error)
      document.getElementById('join-flow').style.display = 'block'
  })
}

function joinSession() {
  try {
    uitoolkit.joinSession(sessionContainer, config)
    uitoolkit.onSessionClosed(sessionClosed)
  } catch (error) {
    console.error('Failed to join session:', error)
    document.getElementById('join-flow').style.display = 'block'
  }
}

var sessionClosed = (() => {
  console.log('session closed')
  if (sessionContainer) {
    uitoolkit.closeSession(sessionContainer)
  }
  document.getElementById('join-flow').style.display = 'block'
})
</script>

<template>
  <main>
    <div id="join-flow">
      <h1>Zoom Video SDK Sample Vue.js</h1>
      <p>User interface offered by the Video SDK UI Toolkit</p>

      <button @click="getVideoSDKJWT">Join Session</button>
    </div>

    <div id='sessionContainer'></div>
  </main>
</template>

<style scoped>
#sessionContainer {
  width: 100%;
  height: 100%;
  min-height: 600px;
}

#join-flow {
  text-align: center;
  padding: 20px;
}
</style>