<template>
  <div>
    <ul v-if="emails.length">
      <li v-for="email in emails" :key="email._id">
        {{ email._source.content }}
      </li>
    </ul>
    <p v-else>No se encontraron emails.</p>
  </div>
</template>

<script>
export default {
  props: {
    emails: {
      type: Array,
      default: () => []
    }
  }
};
</script>

<style scoped>
/* Puedes agregar estilos específicos aquí si lo deseas */
</style>
