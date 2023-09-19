<template>
  <div>
    <!-- Campo de búsqueda y botón -->
    <div class="flex items-center border rounded-full w-3/4 mx-auto mt-4">
      <input v-model="term" placeholder="Buscar emails..." class="flex-grow py-2 px-4 outline-none rounded-l-full" @keyup.enter="search" />
      <button @click="search" class="px-4 text-gray-500 hover:text-blue-400 focus:outline-none bg-white rounded-r-full">
        <i class="fas fa-search"></i>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      term: ''
    };
  },
  methods: {
    extractEmails(content) {
      const regex = /([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+)/gi;
      return content.match(regex) || [];
    },

    async search() {
      if (!this.term) {
        this.$emit('searched', []);
        return;
      }

      try {
        const response = await fetch(`http://localhost:8080/api/search?term=${this.term}`);
        const data = await response.json();

        const emailContents = data.hits.hits.map(hit => ({
          id: hit._id,
          content: hit._source.content,
          emails: this.extractEmails(hit._source.content)
        }));

        this.$emit('searched', emailContents);
      } catch (error) {
        console.error("Error fetching search results:", error);
        this.$emit('searched', []);
      }
    }
  }
};
</script>
