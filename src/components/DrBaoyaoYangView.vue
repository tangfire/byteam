<script setup lang="ts">
// Add data logic here
</script>

<template>
  <el-card class="profile-card">
    <template #header>
      <h2>Professor Baoyao Yang</h2>
    </template>

    <div class="profile-content">
      <img
          src="/currentprojects/yangbaoyao.png"
          alt="Dr. Baoyao Yang"
          class="profile-image"
      />
      <div class="profile-description">
        <p>
          Dr. Baoyao Yang is a distinguished young talent from Guangdong University of Technology. He earned his
          bachelor's degree from South China University of Technology in 2014 and completed his doctorate at
          Hong Kong Baptist University in 2018. Dr. Yang has led various prestigious projects, including National
          Natural Science Foundation of China projects, the Youth Science Foundation of China, and numerous
          provincial and ministerial scientific research initiatives. In 2023, he was awarded the second prize
          for Guangdong Province Science and Technology Progress (ranked fourth individually).
        </p>
        <p>
          In recent years, Dr. Yang has published over 30 papers in top-tier international journals such as IEEE
          TIP, TIFS, and TNNLS, as well as leading conferences like AAAI and MICCAI. He has also been granted
          6 national invention patents. His focus is on the integration of industry, academia, and research. He
          has successfully guided students to win numerous awards in national competitions, and many of his
          outstanding undergraduate students have gone on to study at renowned universities.
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
  color: #7d1231;
  font-size: 1.8rem;
  font-weight: bold;
  border-bottom: 2px solid #7d1231;
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
