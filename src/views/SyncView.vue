<template>
  <button @click="execute">押してね</button>
  {{ data.result.message }}
</template>

<script>
import axios from 'axios'
import {reactive} from "vue";

export default {
  name: "SyncView",
  setup() {
    let data = reactive({
      'result': ''
    })

    const hello = () => {
      console.log(123)
    }

    const hello2 = function b() {
      console.log(456)
    }

    const execute = async () => {
      console.log("start...")
      let res = await axios.get("http://localhost:8000/hello")
      console.log(res.data)
      console.log("end...")
      data.result = res.data
    }

    return {
      hello,
      hello2,
      execute,
      data
    }
  }
}
</script>

<style scoped>

</style>
