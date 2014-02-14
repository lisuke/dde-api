/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
	_image "image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func (image *Image) GetImageSize(imageFile string) (w, h int32, err error) {
	// open the image file
	fr, err := os.Open(imageFile)
	if err != nil {
		// logError(err.Error()) // TODO
		return
	}
	defer fr.Close()

	img, _, err := _image.Decode(fr)
	if err != nil {
		// image format not support
		// logError(err.Error()) // TODO
		return
	}

	w = int32(img.Bounds().Max.X)
	h = int32(img.Bounds().Max.Y)
	return
}
