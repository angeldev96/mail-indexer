

<template>
  <div>
    <!-- Tabla -->
    <table class="table-auto w-full">
      <!-- Encabezados de la tabla -->
      <thead>
        <tr>
          <th class="px-4 py-2">Sent by</th>
          <th class="px-4 py-2">Subject</th>
          <th class="px-4 py-2">Preview</th>
          <th class="px-4 py-2">Date</th>
        </tr>
      </thead>
      <tbody>
        <!-- Lista de correos -->
        <tr v-for="emailContent in emailContents" :key="emailContent._id" @click="selectEmail(emailContent)">
          <!-- Nombre de la persona o entidad que envió el correo -->
          <td class="border px-4 py-2">
            {{ extractField(emailContent.content, 'From') }}
          </td>

          <!-- Subject del correo -->
          <td class="border px-4 py-2">
            {{ extractField(emailContent.content, 'Subject') }}
          </td>

          <!-- Primeras 5 palabras del correo -->
          <td class="border px-4 py-2">
            {{ extractContentPreview(emailContent.content) }}
          </td>

          <!-- Fecha que se envió el correo -->
          <td class="border px-4 py-2">
            {{ formatDate(extractField(emailContent.content, 'Date')) }}
          </td>
        </tr>
      </tbody>
      <EmailDetail v-if="selectedEmail" :email="selectedEmail" />

    </table>
  </div>
</template>

<script>
import EmailDetail from './EmailDetail.vue';

export default {
  
  components: {
    EmailDetail
  },
  props: {
    emailContents: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      selectedEmail: null
    };
  },
  methods: {
    extractField(content, field) {
      const regex = new RegExp(`${field}: (.+)\\r\\n`);
      const match = content.match(regex);
      return match ? match[1].trim() : '';
    },
    extractContentPreview(content) {
      const body = content.split('\r\n\r\n')[1];
      return body ? body.split(' ').slice(0, 5).join(' ') + '...' : '';
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      return `${date.getDate().toString().padStart(2, '0')}/${(date.getMonth() + 1).toString().padStart(2, '0')}/${date.getFullYear()}`;
    },
    selectEmail(email) {
      this.selectedEmail = email;
    }
    
  }
};
</script>

<style scoped>

</style>

