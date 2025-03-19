<script setup lang="ts">
import { ref, onMounted } from 'vue'

const props = defineProps<{
  text: string
}>()

const emit = defineEmits<{
  'typing-finished': []
}>()

const displayText = ref('')
const showCursor = ref(true)
const typingFinished = ref(false)

onMounted(() => {
  startTyping()
  startCursorBlink()
})

const startTyping = () => {
  let charIndex = 0
  const typingInterval = setInterval(() => {
    if (charIndex < props.text.length) {
      displayText.value = props.text.substring(0, charIndex + 1)
      charIndex++
    } else {
      clearInterval(typingInterval)
      typingFinished.value = true
      emit('typing-finished')
    }
  }, 100)
}

const startCursorBlink = () => {
  setInterval(() => {
    if (!typingFinished.value) {
      showCursor.value = !showCursor.value
    }
  }, 500)
}
</script>

<template>
  <span class="typewriter-text">
    {{ displayText }}
    <span v-if="!typingFinished && showCursor" class="cursor">|</span>
  </span>
</template>

<style scoped>
.typewriter-text {
  display: inline-block;
  font-family: monospace;
}

.cursor {
  color: #10A37F;
  font-weight: bold;
}
</style>