mkdir -p gateway
nano gateway/auth.go
use aes_gcm::{
    aead::{Aead, KeyInit, OsRng},
    Aes256Gcm
};
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct AccountRecord {
    pub account_id: String,
    pub username: String,
    pub role: String,
}

pub struct SecureVaultStorage {
    cipher: Aes256Gcm,
}

impl SecureVaultStorage {
    pub fn new(master_key: &[u8]) -> Self {
        let key = aes_gcm::Key::<Aes256Gcm>::from_slice(master_key);
        let cipher = Aes256Gcm::new(key);
        Self { cipher }
    }

    pub fn encrypt_account(&self, account: &AccountRecord) -> Result<(Vec<u8>, Vec<u8>), String> {
        let serialized = serde_json::to_vec(account).map_err(|e| e.to_string())?;
        let nonce = aes_gcm::aead::Nonce::<Aes256Gcm>::generate(&mut OsRng);
        let ciphertext = self.cipher.encrypt(&nonce, serialized.as_ref())
            .map_err(|e| e.to_string())?;
        Ok((ciphertext, nonce.to_vec()))
    }
}

