#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <oqs/oqs.h>

int main() {
    // Algorithm identifier for Kyber512
    const char *alg_id = OQS_KEM_alg_kyber_512;

    // Check if the algorithm is enabled
    if (!OQS_KEM_alg_is_enabled(alg_id)) {
        fprintf(stderr, "Kyber KEM is not enabled.\n");
        return EXIT_FAILURE;
    }

    // Create a KEM object
    OQS_KEM *kem = OQS_KEM_new(alg_id);
    if (kem == NULL) {
        fprintf(stderr, "Error creating KEM object.\n");
        return EXIT_FAILURE;
    }

    // Allocate space for keys and ciphertext
    uint8_t *public_key = malloc(kem->length_public_key);
    uint8_t *secret_key = malloc(kem->length_secret_key);
    uint8_t *ciphertext = malloc(kem->length_ciphertext);
    uint8_t *shared_secret_e = malloc(kem->length_shared_secret);
    uint8_t *shared_secret_d = malloc(kem->length_shared_secret);

    if (public_key == NULL || secret_key == NULL || ciphertext == NULL || shared_secret_e == NULL || shared_secret_d == NULL) {
        fprintf(stderr, "Error allocating memory.\n");
        OQS_KEM_free(kem);
        return EXIT_FAILURE;
    }

    // Generate key pair
    if (OQS_SUCCESS != OQS_KEM_keypair(kem, public_key, secret_key)) {
        fprintf(stderr, "Error generating keypair.\n");
        OQS_KEM_free(kem);
        return EXIT_FAILURE;
    }

    // Encapsulate secret
    if (OQS_SUCCESS != OQS_KEM_encaps(kem, ciphertext, shared_secret_e, public_key)) {
        fprintf(stderr, "Error encapsulating secret.\n");
        OQS_KEM_free(kem);
        return EXIT_FAILURE;
    }

    // Decapsulate secret
    if (OQS_SUCCESS != OQS_KEM_decaps(kem, shared_secret_d, ciphertext, secret_key)) {
        fprintf(stderr, "Error decapsulating secret.\n");
        OQS_KEM_free(kem);
        return EXIT_FAILURE;
    }

    // Check if the shared secrets are equal
    if (memcmp(shared_secret_e, shared_secret_d, kem->length_shared_secret) != 0) {
        fprintf(stderr, "Shared secrets are not equal.\n");
        OQS_KEM_free(kem);
        return EXIT_FAILURE;
    } else {
    	printf(shared_secret_e);
    	printf(shared_secret_d);
        printf("Key encapsulation and decapsulation were successful.\n");
    }

    // Cleanup
    free(public_key);
    free(secret_key);
    free(ciphertext);
    free(shared_secret_e);
    free(shared_secret_d);
    OQS_KEM_free(kem);

    return EXIT_SUCCESS;
}

