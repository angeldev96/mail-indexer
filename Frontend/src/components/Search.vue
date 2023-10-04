<template>
  <form @submit.prevent="search">
    <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
    <div class="relative flex justify-center">
      <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
      </div>
      <input type="search" id="default-search" v-model="term"
        class="block w-3/4 my-3 p-3 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="Search ..." required>
      <button type="submit"
        class="text-white mb-1  absolute right-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
    </div>
  </form>
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
        this.$emit('updateTerm', this.term);
        return;
      }

      // Codificación de la URL
      const encodedTerm = encodeURIComponent(this.term);

      try {
        const response = await fetch(`http://localhost:8080/api/search?term=${encodedTerm}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const data = await response.json();


        const emailContents = data.hits.hits.map(hit => ({
          id: hit._id,
          content: hit._source.content,
          emails: this.extractEmails(hit._source.content)
        }));

        this.$emit('searched', { emailContents: emailContents, totalHits: data.hits.total.value });
      } catch (error) {
        console.error("Error fetching search results:", error);
        this.$emit('searched', []);
      }
    }

  }
};
</script>
