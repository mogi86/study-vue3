<template>
  <div>
    {{ message }}
    {{ hoge }}
  </div>
</template>

<script>
import {ref} from "vue"

export default {
  props: {
    hoge: {
      type: String,
      default: ''
    },
    title: String
  },
  setup(props) {
    const message = ref('This is Header!')
    console.log(props.title)

    return {
      message
    }
  }
}
</script>

<style scoped>

</style>
