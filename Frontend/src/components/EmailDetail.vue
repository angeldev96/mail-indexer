<template>
    <div v-if="email">
      <h2>{{ extractField(email.content, 'Subject') }}</h2>
      <p><strong>Date:</strong> {{ formatDate(extractField(email.content, 'Date')) }}</p>
      <p><strong>From:</strong> {{ extractField(email.content, 'From') }}</p>
      <p><strong>To:</strong> {{ extractField(email.content, 'To') }}</p>
      <p><strong>Content:</strong></p>
      <pre>{{ extractBody(email.content) }}</pre>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      email: {
        type: Object,
        required: true
      }
    },
    methods: {
      extractField(content, field) {
        const regex = new RegExp(`${field}: (.+)\\r\\n`);
        const match = content.match(regex);
        return match ? match[1].trim() : '';
      },
      extractBody(content) {
        return content.split('\r\n\r\n')[1] || '';
      },
      formatDate(dateString) {
        const date = new Date(dateString);
        return `${date.getDate().toString().padStart(2, '0')}/${(date.getMonth() + 1).toString().padStart(2, '0')}/${date.getFullYear()}`;
      },
     
    }
  };
  </script>
  
  <style scoped>


</style>
  