<template>
  <button class="btn btn-primary" @click="execute">メッセージ取得</button>
  <div v-if="data.message && data.isStart">{{ data.message }}</div>
  <div v-else-if="!data.message && data.isStart">メッセージ取得中...</div>
</template>

<script>
import axios from 'axios'
import {reactive} from "vue";

export default {
  name: "SyncView",
  setup() {
    let data = reactive({
      'isStart': false,
      'result': '',
      'message': '',
    })

    const hello = () => {
      console.log(123)
    }

    const hello2 = function b() {
      console.log(456)
    }

    const execute = async () => {
      data.isStart = true
      data.message = ''
      console.log("start...")

      let res = await axios.get("http://localhost:8000/hello")

      console.log(res.data)
      console.log("end...")
      data.message = res.data.message
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

<style scoped lang="scss">
.btn-primary {
  background-color: aqua;
}
</style>
