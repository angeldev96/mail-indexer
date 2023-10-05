<template>
  <div class="flex flex-col items-center p-3">
    <!-- Help text -->
    <span class="text-sm text-gray-700 dark:text-gray-400">
      Showing
      <span class="font-semibold text-gray-900 dark:text-white">{{ (currentPage - 1) * entriesPerPage + 1 }}</span>
      to
      <span class="font-semibold text-gray-900 dark:text-white">{{ Math.min(currentPage * entriesPerPage, totalHits) }}</span>
      of
      <span class="font-semibold text-gray-900 dark:text-white">{{ totalHits }}</span>
      Entries
    </span>
    <!-- Pagination buttons -->
    <div class="xs:mt-0 mt-2 inline-flex">
      <!-- Buttons -->
      <button 
        :disabled="currentPage === 1"
        @click="prevPage" 
        class="flex h-8 items-center justify-center rounded-l bg-gray-800 px-3 text-sm font-medium text-white hover:bg-gray-900 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
      >
        <svg class="mr-2 h-3.5 w-3.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5H1m0 0 4 4M1 5l4-4" />
        </svg>
        Prev
      </button>
      <button 
        @click="nextPage" 
        class="flex h-8 items-center justify-center rounded-r border-0 border-l border-gray-700 bg-gray-800 px-3 text-sm font-medium text-white hover:bg-gray-900 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
      >
        Next
        <svg class="ml-2 h-3.5 w-3.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    totalHits: {
      type: Number,
      required: true
    },
    entriesPerPage: {
      type: Number,
      default: 10
    }
  },
  data() {
    return {
      currentPage: 1
    };
  },
  methods: {
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.$emit('pageChanged', this.currentPage);
      }
    },

    nextPage() {
      if (this.currentPage < Math.ceil(this.totalHits / this.entriesPerPage)) {
        this.currentPage++;
        this.$emit('pageChanged', this.currentPage);
      }
    }

  }
};
</script>
