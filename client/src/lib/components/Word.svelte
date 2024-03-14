<script lang="ts">
  let checked: boolean = false;
</script>
<button on:click={() => checked = !checked} class="flex p-2 items-center justify-center grow border-black border-2 rounded " class:bg-gray-400={checked}>

  <p><slot /></p>
</button>
