<script setup lang="ts">
// Add data logic here
</script>

<template>
  <el-card class="profile-card">
    <template #header>
      <h2>Dr. Baoyao Yang(Associate Professor)</h2>
    </template>

    <div class="profile-content">
      <img
          src="/currentprojects/yangbaoyao.jpg"
          alt="Dr. Baoyao Yang"
          class="profile-image"
      />
      <div class="profile-description">
        <p>
          Dr. Baoyao Yang, a distinguished young faculty member recruited under the "Youth Hundred Talents Program" at Guangdong University of Technology, obtained her Bachelor’s degree from South China University of Technology in 2014 and her Ph.D. from Hong Kong Baptist University in 2018. She has spearheaded several key research projects, including the National Natural Science Foundation General Program, Youth Science Fund, and multiple provincial/ministerial-level scientific research initiatives. Her contributions were recognized with the 2023 Guangdong Province Science and Technology Progress Award (Second Prize, ranked fourth individually).

        </p>
        <p>
          Dr. Yang has authored over 30 publications in top-tier journals such as IEEE TIP, TIFS, and TNNLS, as well as leading conferences like AAAI and MICCAI. She holds six authorized national invention patents and emphasizes the integration of industry, academia, and research. Committed to mentorship, she has guided students to win awards in national competitions and helped numerous undergraduates gain admission to prestigious institutions for advanced studies.

        </p>
      </div>
    </div>
  </el-card>
</template>

<style scoped>
.profile-card {
  max-width: 1000px;
  margin: 2rem auto;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.profile-card h2 {
  color: #2c3e50;
  font-size: 1.8rem;
  font-weight: bold;
  /*border-bottom: 2px solid #7d1231;*/
  padding-bottom: 0.5em;
  margin-bottom: 0.5em;
}

.profile-content {
  display: flex;
  flex-wrap: wrap;
  gap: 2rem;
  align-items: flex-start;
}

.profile-image {
  width: 300px;
  height: auto;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.profile-description {
  flex: 1;
  font-family: 'Segoe UI', Roboto, sans-serif;
  line-height: 1.8;
  color: #2d3436;
  font-size: 1rem;
  text-align: justify;  /* 添加两端对齐 */
  hyphens: auto;        /* 允许自动断字 */
  word-break: break-word; /* 长单词强制换行 */
}

.profile-description p {
  margin-bottom: 1.2em;
}

.profile-description p:last-child {
  margin-bottom: 0;
}

@media (max-width: 768px) {
  .profile-content {
    flex-direction: column;
    align-items: center;
  }

  .profile-image {
    width: 100%;
    max-width: 400px;
  }
}
</style>
