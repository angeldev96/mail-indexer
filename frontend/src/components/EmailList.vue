<template>
  <div>
    <!-- Lista de correos y su contenido -->
    <div v-for="emailContent in emailContents" :key="emailContent.id" class="mt-4 p-4 border rounded">
      <h3 class="text-xl font-bold mb-2">Contenido del Correo:</h3>
      <p>{{ emailContent.content }}</p>
      <h4 class="text-lg font-bold mt-4 mb-2">Direcciones de correo electrónico:</h4>
      <ul>
        <li v-for="email in emailContent.emails" :key="email">{{ email }}</li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    emailContents: {
      type: Array,
      required: true
    }
  }
};
</script>

<style scoped>
</style>
