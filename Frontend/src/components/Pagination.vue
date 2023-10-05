<template>
  <div class="flex flex-col items-center p-3">
    <!-- Help text -->
    <span class="text-sm text-gray-700 dark:text-gray-400">
      Showing
      <span class="font-semibold text-gray-900 dark:text-white">1</span>
      to
      <span class="font-semibold text-gray-900 dark:text-white">10</span>
      of
      <span class="font-semibold text-gray-900 dark:text-white">{{ totalHits }}</span>
      Entries
    </span>
    <!-- ... (el resto del template de paginación) ... -->
  </div>
</template>

<script>
export default {
  props: {
    totalHits: {
      type: Number,
      required: true
    }
  }
};
</script>
