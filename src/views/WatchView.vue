<template>
  <form>
    <input type="text" v-model="textValue">
  </form>
  <p>price: {{ price }}</p>
  <p>price2: {{ price2 }}</p>
</template>

<script>
import {ref, watch} from "vue"

export default {
  name: "WatchView",
  setup() {
    const textValue = ref(0)
    let price = ref(0)
    let price2 = 0

    watch(textValue, (newValue) => {
      console.log("newValue: ", newValue)
      price.value = newValue * 1.10
      price2 = newValue * 1.10 // It is not changed because don't use `ref()`
    })

    return {
      textValue,
      price,
      price2
    }
  }
}
</script>

<style scoped>

</style>
