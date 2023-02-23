// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import  { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyAiwUOodibl1ouMsHPSiC1jbF76M5TlmBM",
  authDomain: "taskmanager-c997b.firebaseapp.com",
  projectId: "taskmanager-c997b",
  storageBucket: "taskmanager-c997b.appspot.com",
  messagingSenderId: "695268071989",
  appId: "1:695268071989:web:cc5e26508948339c9aef7c",
  measurementId: "G-QZTMN3HSFJ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

// Initialize Firebase Authentication and get a reference to the service
const auth = getAuth(app);
