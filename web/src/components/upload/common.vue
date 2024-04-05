<template>
  <div>
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload?id=${props.id}&type=${props.type}`"
      :before-upload="checkFile"
      :data="uploadData"
      :headers="{ 'x-token': userStore.token }"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button type="primary">上传配置文件</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { isImageMime, isVideoMime } from '@/utils/image'

defineOptions({
  name: 'UploadCommon',
})

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)
const props = defineProps({
  id: {
    default: 0
  },
  type: {
    default: ''
  },
})
const userStore = useUserStore()
const fullscreenLoading = ref(false)
let uploadData = {}
const checkFile = (file) => {
  uploadData = {
    'id': props.id,
    'type': props.type
  }
  fullscreenLoading.value = true
  const isLt500K = file.size / 1024 / 1024 < 0.5 // 500K, @todo 应支持在项目中设置
  const isLt5M = file.size / 1024 / 1024 < 5 // 5MB, @todo 应支持项目中设置
  const isVideo = isVideoMime(file.type)
  const isImage = isImageMime(file.type)
  let pass = true
  // if (!isVideo && !isImage) {
  //   ElMessage.error('上传图片只能是 jpg,png,svg,webp 格式, 上传视频只能是 mp4,webm 格式!')
  //   fullscreenLoading.value = false
  //   pass = false
  // }
  // if (!isLt5M && isVideo) {
  //   ElMessage.error('上传视频大小不能超过 5MB')
  //   fullscreenLoading.value = false
  //   pass = false
  // }
  // if (!isLt500K && isImage) {
  //   ElMessage.error('未压缩的上传图片大小不能超过 500KB，请使用压缩上传')
  //   fullscreenLoading.value = false
  //   pass = false
  // }
  if (!isLt500K) {
    ElMessage.error('上传文件不能超过500K!')
    fullscreenLoading.value = false
    pass = false
  }

  console.log('upload file check result: ', pass)

  return pass
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

</script>

