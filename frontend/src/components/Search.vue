<template>
  <div>
    <!-- Campo de búsqueda y botón -->
    <div class="flex items-center border rounded-full w-3/4 mx-auto mt-4">
      <input v-model="term" placeholder="Buscar emails..." class="flex-grow py-2 px-4 outline-none rounded-l-full" @keyup.enter="search" />
      <button @click="search" class="px-4 text-gray-500 hover:text-blue-400 focus:outline-none bg-white rounded-r-full">
        <i class="fas fa-search"></i>
      </button>
    </div>

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
  data() {
    return {
      term: '',
      emailContents: []
    };
  },
  methods: {
    extractEmails(content) {
      const regex = /([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+)/gi;
      return content.match(regex) || [];
    },

    async search() {
      if (!this.term) {
        this.emailContents = [];
        return;
      }

      try {
        const response = await fetch(`http://localhost:8080/api/search?term=${this.term}`);
        const data = await response.json();

        this.emailContents = data.hits.hits.map(hit => ({
          id: hit._id,
          content: hit._source.content,
          emails: this.extractEmails(hit._source.content)
        }));
      } catch (error) {
        console.error("Error fetching search results:", error);
        this.emailContents = [];
      }
    }
  }
};
</script>

<style scoped>
/* Aquí puedes agregar estilos específicos para este componente si lo necesitas */
</style>
