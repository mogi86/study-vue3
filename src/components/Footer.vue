<template>
  <div>
<!--    {{ message }}-->
    {{ message2 }}
  </div>
</template>

<script>
import {ref} from "vue"

// const message = ref('This is Footer!')

export default {
  setup() {
    const message2 = ref('This is Footer2!')
    return {
      message2
    }
  }
}
</script>

<style scoped>

</style>
