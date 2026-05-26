import type { App, Plugin } from 'vue'
import {
  ElAlert,
  ElAside,
  ElColorPicker,
  ElDatePicker,
  ElDialog,
  ElForm,
  ElFormItem,
  ElImage,
  ElInput,
  ElInputNumber,
  ElLoading,
  ElOption,
  ElPagination,
  ElPopconfirm,
  ElSegmented,
  ElSelect,
  ElSwitch,
  ElTable,
  ElTableColumn,
  ElTag,
  ElUpload,
} from 'element-plus'

import 'element-plus/es/components/alert/style/css'
import 'element-plus/es/components/aside/style/css'
import 'element-plus/es/components/color-picker/style/css'
import 'element-plus/es/components/date-picker/style/css'
import 'element-plus/es/components/dialog/style/css'
import 'element-plus/es/components/form/style/css'
import 'element-plus/es/components/image/style/css'
import 'element-plus/es/components/input/style/css'
import 'element-plus/es/components/input-number/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/message-box/style/css'
import 'element-plus/es/components/option/style/css'
import 'element-plus/es/components/pagination/style/css'
import 'element-plus/es/components/popconfirm/style/css'
import 'element-plus/es/components/segmented/style/css'
import 'element-plus/es/components/select/style/css'
import 'element-plus/es/components/switch/style/css'
import 'element-plus/es/components/table/style/css'
import 'element-plus/es/components/table-column/style/css'
import 'element-plus/es/components/tag/style/css'
import 'element-plus/es/components/upload/style/css'

const adminComponents: Plugin[] = [
  ElAlert,
  ElAside,
  ElColorPicker,
  ElDatePicker,
  ElDialog,
  ElForm,
  ElFormItem,
  ElImage,
  ElInput,
  ElInputNumber,
  ElOption,
  ElPagination,
  ElPopconfirm,
  ElSegmented,
  ElSelect,
  ElSwitch,
  ElTable,
  ElTableColumn,
  ElTag,
  ElUpload,
]

let installed = false

export const installAdminElementPlus = (app: App) => {
  if (installed) return
  adminComponents.forEach((component) => app.use(component))
  app.use(ElLoading)
  installed = true
}
