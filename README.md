# Pi Dashboard Setup Guide

This guide explains how to set up the Pi Dashboard on a Raspberry Pi.

---

## Mounting External Drives

To mount an external drive (e.g., Toshiba):

1. Edit the `fstab` file:
   ```bash
   sudo nano /etc/fstab
   ```
2. Add the following line (replace the UUID with your drive’s UUID and the path with a new path):
   ```bash
   UUID=7299-3935 /mnt/toshiba-white exfat defaults,auto,users,rw,nofail,uid=1000,gid=126,noatime 0 0
   ```

