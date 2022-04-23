<template>
  <div class="no-parallel">
    <button type="button" class="btn btn-primary" @click="noParallel">直列で実行</button>
  </div>
  <div class="parallel">
    <button type="button" class="btn btn-primary" @click="parallel">並行で実行</button>
  </div>
</template>

<script>

import axios from "axios";
import {reactive} from "vue";

export default {
  name: "SyncView2",
  setup() {
    let data = reactive({
      'isStart': false,
      'result': '',
      'message': '',
    })

    const execute = async () => {
      data.isStart = true
      data.message = ''
      console.log("start...")
      let res = await axios.get("http://localhost:8000/hello")
      console.log("end...")
      return res
    }

    const execute2 = async () => {
      data.isStart = true
      data.message = ''
      console.log("start2...")
      let res = await axios.get("http://localhost:8000/hello2")
      console.log("end2...")
      return res
    }

    const parallel = async function fn() {
      await Promise.all([
        execute(),
        execute2()
      ])
    }

    // It seems that `noParallel` also execute two calling api in parallel
    const noParallel = function fn() {
      execute()
      execute2()
    }

    return {
      execute,
      parallel,
      noParallel,
    }
  }
}
</script>

<style scoped lang="scss">
.no-parallel {
  padding-bottom: 10px;
}
</style>
