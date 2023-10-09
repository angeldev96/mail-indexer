<template>
  <div style="border-bottom: 1px solid rgb(55, 65, 81); border-top-color: rgb(55, 65, 81); border-right-color: rgb(55, 65, 81); border-left-color: rgb(55, 65, 81); border-top-style: solid; border-right-style: solid; border-left-style: solid;" class="items-center bg-gray-800 flex justify-between p-4">
  <div class="items-center flex">
    <div class="pr-3">
      <a href="https://flowbite.com/application-ui/demo/mailing/inbox/" class="text-gray-400 cursor-pointer inline-flex justify-center rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path fill-rule="evenodd" d="M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
    </div>
    <div class="border-l flex px-2 border-gray-700 border-solid">
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer flex justify-center rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer flex justify-center ml-2 rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer flex justify-center ml-2 rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6">
          <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z" fill="rgb(156, 163, 175)"></path>
          <path fill-rule="evenodd" d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path>
        </svg>
      </a>
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer flex justify-center ml-2 rounded p-1">
        <svg fill="none" stroke="#9ca3af" viewbox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="0.13rem" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" fill="none" stroke="#9ca3af"></path></svg>
      </a>
    </div>
    <div class="border-l text-gray-500 text-sm font-medium pl-3 border-gray-700 border-solid"> Date: {{ formatDate(extractField(email.content, 'Date')) }}</div>
  </div>
  <div class="flex px-2">
    <div class="pr-2">
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer inline-flex justify-center rounded p-1">
        <svg fill="none" stroke="#9ca3af" viewbox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="0.13rem" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" fill="none" stroke="#9ca3af"></path></svg>
      </a>
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer inline-flex justify-center rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6"><path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
    </div>
    <div class="border-l ml-2 pl-2 border-gray-700 border-solid">
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer inline-flex justify-center rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-7 w-7"><path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
      <a href="https://flowbite.com/#" class="text-gray-400 cursor-pointer inline-flex justify-center rounded p-1">
        <svg fill="rgb(156, 163, 175)" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" class="h-7 w-7"><path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" fill="rgb(156, 163, 175)"></path></svg>
      </a>
    </div>
  </div>
</div>

  <div v-if="email" class="bg-gray-900 p-5">
  
    <!-- Remitente y Avatar -->
    <div class="items-center flex mb-4">
      <div>
        <!-- Puedes cambiar la URL de la imagen si tienes una dinámica basada en el remitente -->
        <img src="https://flowbite.com/application-ui/demo/images/users/bonnie-green.png" class="h-8 w-8 align-middle rounded-full overflow-clip" />
      </div>
      <div class="ml-4">
        <div class="text-white font-semibold text-ellipsis overflow-hidden">{{ extractField(email.content, 'From') }}</div>
      </div>
    </div>

    <!-- Asunto del correo -->
    <h1 class="text-white text-2xl font-bold mb-4">{{ extractField(email.content, 'Subject') }}</h1>

    <!-- Contenido del correo -->
    <div>
      <pre class="text-gray-400">{{ extractBody(email.content) }}</pre>
    </div>

  </div>
</template>

  
  <script>
  export default {
    props: {
      email: {
        type: Object,
        required: true
      },
      term: {
      type: String,
      default: ''
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
  