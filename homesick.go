package homesick

import (
  "os"
  "path/filepath"
  "runtime"
)

func Cache(application string) string {
  switch runtime.GOOS {
  case "darwin":
    return filepath.Join(home(), "Library", "Caches", application)
  case "windows":
    return appData(application)
  default:
    return filepath.Join(xdgCache(), application)
  }
}

func xdgCache() string {
  xdg := os.Getenv("XDG_CACHE_HOME")
  if xdg == "" {
    return filepath.Join(home(), ".cache")
  }
  return xdg
}


func Config(application string) string {
  switch runtime.GOOS {
  case "darwin":
    return filepath.Join(home(), "Library", "Application Support", application)
  case "windows":
    return appData(application)
  default:
    return filepath.Join(xdgConfig(), application)
  }
}

func xdgConfig() string {
  xdg := os.Getenv("XDG_CONFIG_HOME")
  if xdg == "" {
    return filepath.Join(home(), ".config")
  }
  return xdg
}

func appData(application string) string {
  return filepath.Join(os.Getenv("APPDATA"), application)
}

func home() string {
  return os.Getenv("HOME")
}