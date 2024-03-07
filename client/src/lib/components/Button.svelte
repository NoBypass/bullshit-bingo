<script lang="ts">
    import { createEventDispatcher } from "svelte";

  export const text: string = "button";

  const dispatch = createEventDispatcher();
</script>

<button class="p-1 hover:bg-gray-100 border-black border-2 rounded" on:click={() => dispatch('click')}>
  <slot />
</button>
