use aes_gcm::{
    aead::{Aead, KeyInit, OsRng},
    Aes256Gcm, Nonce, Key
};
use rand::RngCore;

pub struct CryptoEngine {
    key: Key<Aes256Gcm>,
}

impl CryptoEngine {
    /// تهيئة المحرك بمفتاح يتم استقباله مسبقاً (من متغيرات البيئة أو الخزنة)
    pub fn new(secret_key: &[u8]) -> Result<Self, &'static str> {
        if secret_key.len() != 32 {
            return Err("Invalid key length: AES-256 requires exactly 32 bytes");
        }
        let key = *Key::<Aes256Gcm>::from_slice(secret_key);
        Ok(Self { key })
    }

    /// دالة مساعدة لتوليد مفتاح جديد (تستخدم فقط عند التهيئة الأولى للإعدادات وليس عند تشغيل الخادم)
    pub fn generate_master_key() -> Vec<u8> {
        let key = Aes256Gcm::generate_key(&mut OsRng);
        key.to_vec()
    }

    pub fn encrypt(&self, data: &[u8]) -> Result<(Vec<u8>, Vec<u8>), &'static str> {
        let cipher = Aes256Gcm::new(&self.key);
        
        // توليد Nonce فريد بطول 12 بايت لكل عملية تشفير
        let mut nonce_bytes = [0u8; 12];
        OsRng.fill_bytes(&mut nonce_bytes);
        let nonce = Nonce::from_slice(&nonce_bytes);

        let ciphertext = cipher.encrypt(nonce, data)
            .map_err(|_| "Encryption failed")?;

        Ok((ciphertext, nonce_bytes.to_vec()))
    }

    pub fn decrypt(&self, ciphertext: &[u8], nonce_bytes: &[u8]) -> Result<Vec<u8>, &'static str> {
        let cipher = Aes256Gcm::new(&self.key);
        let nonce = Nonce::from_slice(nonce_bytes);

        let plaintext = cipher.decrypt(nonce, ciphertext)
            .map_err(|_| "Decryption failed")?;

        Ok(plaintext)
    }
}

