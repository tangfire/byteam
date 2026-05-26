import type { App, Plugin } from 'vue'
import {
  ElAvatar,
  ElBacktop,
  ElButton,
  ElCard,
  ElCollapseTransition,
  ElContainer,
  ElEmpty,
  ElFooter,
  ElHeader,
  ElIcon,
  ElMain,
  ElMenu,
  ElMenuItem,
  ElSpace,
  ElSubMenu,
  ElTimeline,
  ElTimelineItem,
} from 'element-plus'

import 'element-plus/es/components/avatar/style/css'
import 'element-plus/es/components/backtop/style/css'
import 'element-plus/es/components/button/style/css'
import 'element-plus/es/components/card/style/css'
import 'element-plus/es/components/container/style/css'
import 'element-plus/es/components/empty/style/css'
import 'element-plus/es/components/icon/style/css'
import 'element-plus/es/components/menu/style/css'
import 'element-plus/es/components/space/style/css'
import 'element-plus/es/components/timeline/style/css'

const components: Plugin[] = [
  ElAvatar,
  ElBacktop,
  ElButton,
  ElCard,
  ElCollapseTransition,
  ElContainer,
  ElEmpty,
  ElFooter,
  ElHeader,
  ElIcon,
  ElMain,
  ElMenu,
  ElMenuItem,
  ElSpace,
  ElSubMenu,
  ElTimeline,
  ElTimelineItem,
]

export const installElementPlus = (app: App) => {
  components.forEach((component) => app.use(component))
}
