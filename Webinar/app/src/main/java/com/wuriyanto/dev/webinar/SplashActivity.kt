package com.wuriyanto.dev.webinar

import android.content.Intent
import android.os.Bundle
import android.os.Handler
import android.os.Looper
import androidx.appcompat.app.AppCompatActivity

class SplashActivity : AppCompatActivity() {

    private val SPLASH_TIMEOUT: Long = 3000
    private lateinit var handler: Handler
    private lateinit var callback: Runnable

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_splash)

        handler = Handler(Looper.getMainLooper())
        callback = Runnable {
            this@SplashActivity.startActivity(Intent(this@SplashActivity, SignInActivity::class.java))
            finish()
        }
        handler.postDelayed(callback, SPLASH_TIMEOUT)
    }

    override fun onDestroy() {
        super.onDestroy()
        if (handler != null) {
            handler.removeCallbacks(callback)
        }
    }
}