import { deleteUser } from '../services/userService.js';

export function attachDeleteEvents() {
    document.querySelectorAll('.delete').forEach(btn => {
        btn.addEventListener('click', async function () {
            try {
                await deleteUser(this.getAttribute('data-id'));
                window.location.reload();
            } catch (error) {
                console.error(error);
                alert('Failed to delete user.');
            }
        });
    });
}
