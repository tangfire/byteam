<script setup lang="ts">
import { ref } from 'vue'

const currentYear = ref(new Date().getFullYear())
</script>

<template>
  <el-footer :height="120" class="custom-footer">
    <div class="footer-content">
      <p class="footer-text">
        Welcome to BYML @ 
        <a 
          class="gdutlink" 
          href="https://www.gdut.edu.cn/"
          target="_blank"
        >
          Guangdong University of Technology
        </a>
      </p>

      <p class="footer-text">
        Email: ybaoyao@gdut.edu.cn
      </p>

      <a 
        class="BaoyaoGroupLink" 
        href="https://github.com/BaoyaoGroup"
        target="_blank"
      >
        Github-BYML
      </a>

      <p class="copyright">
        © {{ currentYear }} By Baoyao Yang.
      </p>
    </div>
  </el-footer>
</template>

<style scoped>
.custom-footer {
  background-color: var(--bg-dark);
  color: white;
  padding: 20px 0;
}

.footer-content {
  text-align: center;
  padding: 20px;
  color: #fff;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 8px;
}

.footer-text {
  margin: 5px 0;
  font-size: var(--font-size-sm);
}

.copyright {
  margin-top: 20px;
  font-size: var(--font-size-sm);
  opacity: 0.9;
}

.gdutlink,
.BaoyaoGroupLink {
  text-decoration: none;
  color: white;
  transition: color var(--transition-base);
  
  &:hover {
    color: #747d8c;
  }
}

@media (max-width: 768px) {
  .custom-footer {
    padding: 15px 0;
    height: auto !important;
  }
  
  .footer-content {
    gap: 12px;
  }
  
  .footer-text {
    font-size: var(--font-size-xs);
  }
}
</style>
