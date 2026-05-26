import { getCurrentInstance } from 'vue'
import { installAdminElementPlus } from '../plugins/element-plus-admin'

export const useAdminElementPlus = () => {
  const app = getCurrentInstance()?.appContext.app
  if (app) installAdminElementPlus(app)
}
