<!-- Rce.vue -->
<template>
  <div class="rce-app">
    <div>
      <label for="cmd">命令：</label>
      <input
          id="cmd"
          v-model="command"
          class="cmd-input"
          type="text"
      >
    </div>

    <button @click="cmd" class="rce-btn">执行命令</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {ToolService} from "../../bindings/changeme/service";
const props = defineProps({
  targetUrl: {
    type: String,
    default: ''
  },
  selectPoc: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['execute-cmd'])

const command = ref('')

const cmd = () => {
  var action = "exploit"
  var expType = "cmd"
  ToolService.POC(props.targetUrl,props.selectPoc,action,command.value,"",expType).then((result)=>{
    emit('execute-cmd', result)
  })

}


</script>

<style>
.rce-app{
  display: flex;
  padding: 8px 20px;
}
.cmd-input {
  padding: 8px 12px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  font-size: 14px;
  width: 550px;
  margin-left: 10px;
}
.rce-btn {
  background: #f5f5f5;
  cursor: pointer;
  white-space: nowrap;
  padding: 8px 15px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  font-size: 14px;
  margin-left: 15px;

}

</style>