<template>
  <div class="upload-component">
    <!-- 文件名输入框 -->
    文件名：
    <a-input
        id="req_filename"
        v-model="fileName"
        placeholder="1.txt"
        style="flex:1; margin-right: 10px; background: #ffffff;"
    />

    <!-- 文件选择按钮 -->
    <input
        type="file"
        @change="handleFileChange"
        ref="fileInput"
        style="display: none"
    />
    <a-button
        type="primary"
        @click="$refs.fileInput.click()"
        class="action-btn"
    >
      选择文件
    </a-button>

    <!-- 上传按钮 -->
    <a-button
        @click="handleUpload"
        class="action-btn"
        style="margin-left: 10px;"
        type="primary"
    >
      上传文件
    </a-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {ToolService} from "../../bindings/changeme/service";
const emit = defineEmits(['file-upload'])
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

const fileName = ref('')
const fileContent = ref('')
const fileInput = ref(null)

const handleFileChange = (event) => {
  const files = event.target.files
  if (files && files.length > 0) {
    const file = files[0]
    fileName.value = file.name // 自动填充文件名

    const reader = new FileReader()

    reader.onload = (e) => {
      fileContent.value = e.target.result
    }

    reader.onerror = () => {
      console.error('文件读取失败')
    }

    reader.readAsText(file)
  }
}

const handleUpload = () => {
  if (!fileContent.value) {
    alert('请先选择文件')
    return
  }

  // ✅ 正确获取用户输入的文件名
  const name = fileName.value
  var action = "exploit"
  ToolService.POC(props.targetUrl, props.selectPoc, action,fileContent.value,name).then((result)=>{
    emit('file-upload', result)
  })
}
</script>

<style scoped>
.upload-component {
  display: flex;
  align-items: center;
  padding: 5px 25px;
}

/* 按钮统一样式 */
.action-btn {
  height: 32px;
  font-size: 14px;
  padding: 0 12px;
  white-space: nowrap;
  width: 130px;
}
</style> 