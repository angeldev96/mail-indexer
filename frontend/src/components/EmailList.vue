<!-- src/components/EmailList.vue -->

<template>
    <div>
      <ul v-if="emails && emails.length">
        <li v-for="email in emails" :key="email._id" class="border-b p-4">
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
    },
    mounted() {
      console.log("Emails recibidos en EmailList:", this.emails);
    }
  };
  </script>
  
  