/*
 * Poly2Tri Copyright (c) 2009-2011, Poly2Tri Contributors
 * http://code.google.com/p/poly2tri/
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 * * Neither the name of Poly2Tri nor the names of its contributors may be
 *   used to endorse or promote products derived from this software without specific
 *   prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
package p2t

/**
 * Sweep-line, Constrained Delauney Triangulation (CDT) See: Domiter, V. and
 * Zalik, B.(2008)'Sweep-line algorithm for constrained Delaunay triangulation',
 * International Journal of Geographical Information Science
 *
 * "FlipScan" Constrained Edge Algorithm invented by Thomas Åhlén, thahlen@gmail.com
 */

import (
	"fmt"
)

var tcx *SweepContext

func Init(polyline PointArray) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("p2t.Init panicked with: %#v", e)
		}
	}()

	tcx = new(SweepContext)
	tcx.init(polyline)
	return err
}

// Returns the contstrained triangles
func Triangulate() (triangles TriArray, err error) {
	if tcx != nil {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("p2t.Triangle panicked with: %#v", e)
			}
		}()
		triangulate(tcx)
	} else {
		panic(fmt.Sprintf("ERROR: p2t uninitialized"))
	}
	// Copy triangles from list to slice
	triangles = make(TriArray, tcx.triangles.Len())
	i := 0
	for e := tcx.triangles.Front(); e != nil; e = e.Next() {
		triangles[i] = e.Value.(*Triangle)
		i++
	}
	return triangles, err
}

func AddHole(polyline PointArray) {
	if tcx != nil {
		tcx.addHole(polyline)
	} else {
		panic(fmt.Sprintf("ERROR: p2t uninitialized"))
	}
}

func AddPoint(p *Point) {
	if tcx != nil {
		tcx.addPoint(p)
	} else {
		panic(fmt.Sprintf("ERROR: p2t uninitialized"))
	}
}

// Returns the entire triangle mesh for debugging purposes
func Mesh() TriArray {
	if tcx != nil {
		// Convert from Vector to slice for convenience
		n := tcx.tmap.Len()
		var triangles = make(TriArray, n)
		for e, i := tcx.tmap.Front(), 0; e != nil; e, i = e.Next(), i+1 {
			triangles[i] = e.Value.(*Triangle)
		}
		return triangles
	}
	panic(fmt.Sprintf("ERROR: p2t uninitialized"))
}
