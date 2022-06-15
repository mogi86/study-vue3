<template>
  <Header
      hoge="hogehogeeeee"
      class="test"
      title="titledayo"
      :fuga="fugafuga"
  ></Header>
  <div>
    {{ message }}
  </div>
  <Footer/>
</template>

<script setup>
import Header from "@/components/Header"
import Footer from "@/components/Footer"
import {ref} from "vue"

const message = ref('hello')

</script>

<style scoped>

</style>
